#!/usr/bin/ruby -w

gem 'ruby-cute', ">=0.6"
require 'cute'
require 'pp' # pretty print
require 'distem'
require_relative 'conf'

SITE_NAME = "nancy"
HOSTNAME = "g5k"

g5k = Cute::G5K::API.new(:username => "ksonbol")
jobs = g5k.get_my_jobs(SITE_NAME)
raise "No jobs running! Run ruby platform_setup.rb --reserve to create a job" unless jobs.length() > 0

if jobs.length() > 1
  puts "WARNING: You have multiple jobs running at #{SITE_NAME}"
end

job = jobs.first
subnet = g5k.get_subnets(job).first
subnet_addr = "#{subnet.address}/#{subnet.prefix}"

pnodes = job['assigned_nodes']
pnodes.map!{|n| n.split(".")[0]}  # remove the ".SITE_NAME.grid5000.fr" suffix
raise 'This experiment requires at least two physical machines' unless pnodes.size >= 2
coordinator = pnodes.first

instance_name = "edge"
cluster_token = "edge-cluster"

# these values works fine for both cloud and edge (for now)
hb_interval = 10   # heartbeat interval in ms
elec_timeout = 100 # election timeout in ms

# case SETUP
# when "cloud"
#     hb_interval = 10   # heartbeat interval in ms
#     elec_timeout = 100 # election timeout in ms
# when "edge"
#     hb_interval = 15   # heartbeat interval (1.5xRTT) in ms
#     elec_timeout = 150 # election timeout in ms
# end

serv_node_ips = Array.new(NUM_SERVERS)
initial_cluster_str = "" # needed for etcd peers (servers)
Distem.client do |dis|
    # get node IPs and prepare initial cluster conf
    SERVER_VNODES.each_with_index do |node, idx|
        addr = dis.viface_info(node,'if0')['address'].split('/')[0]
        serv_node_ips[idx] = addr
        initial_cluster_str += "#{node}=http://#{addr}:2380,"
        # if this is already in the fs, no need to export ETCDCTL_API=3
        dis.vnode_execute(node, "pkill etcd")  # kill any previous instances of etcd
    end
    initial_cluster_str = initial_cluster_str[0..-2]  # remove the last comma
    sleep(5)  # make sure old etcd instances are dead
    SERVER_VNODES.each_with_index do |node, idx|
        # clean the log folder
        dis.vnode_execute(node, "rm -rf /root/etcdlog /root/#{node}.etcd; mkdir /root/etcdlog") 
        # dis.vnode_execute(node, "rm -rf /root/#{node}.etcd") 
        addr = serv_node_ips[idx]
        # puts dis.vnode_execute(node, "etcd --version")
        cmd =  "nohup /usr/local/bin/etcd --heartbeat-interval=#{hb_interval} \
        --election-timeout=#{elec_timeout} \
        --name #{node} --initial-advertise-peer-urls http://#{addr}:2380 \
        --listen-peer-urls http://#{addr}:2380 \
        --listen-client-urls http://#{addr}:2379,http://127.0.0.1:2379 \
        --advertise-client-urls http://#{addr}:2379 \
        --initial-cluster-token #{cluster_token} \
        --initial-cluster #{initial_cluster_str} \
        --initial-cluster-state new > /root/etcdlog/etcd.log 2>&1 &"
        # --initial-cluster-state new &> /root/etcdlog/out.log &"
        # IMPORTANT: without the last part of the command the function blocks forever!
        # should we add this to listen-client-urls? ,http://127.0.0.1:4001
        puts dis.vnode_execute(node, cmd)
        # puts "etcd server #{idx+1} running"
    end
    sleep(7)
    dis.vnode_execute(SERVER_VNODES[0], "etcdctl put mykey myvalue")
    sleep(3)
    out = dis.vnode_execute(SERVER_VNODES[1], "etcdctl get mykey")
    if out.length>=2 && out[1] == "myvalue"
        puts "etcd cluster is working correctly"
    else
        puts "etcd cluster not setup correctly: '#{out}'"
    end
    dis.vnode_execute(SERVER_VNODES[2], "etcdctl del mykey")
end
# puts "all etcd servers are now running!"