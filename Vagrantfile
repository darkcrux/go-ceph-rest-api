# -*- mode: ruby -*-
# vi: set ft=ruby :

VAGRANTFILE_API_VERSION = "2"

Vagrant.configure(VAGRANTFILE_API_VERSION) do |config|

  HOSTNAME = 'ceph-rest-api'

  config.vm.box = "ubuntu/trusty64"

  config.vm.define HOSTNAME
  config.vm.hostname = HOSTNAME
  config.vm.network :private_network, ip: '200.200.200.200'

  config.vm.provider :virtualbox do |vb|
    2.times.each do |i|
      vb.customize [ "createhd", "--filename", "ceph-disk-#{i}", "--size", "10000" ]
      vb.customize [ "storageattach", :id, "--storagectl", "SATAController", "--port", 3+i, "--device", 0, "--type", "hdd", "--medium", "ceph-disk-#{i}.vdi" ]        
    end    
  end

  config.vm.provision :ansible do |ansible|
    ansible.groups = {
      'mons' => [ HOSTNAME ],
      'osds' => [ HOSTNAME ],
      'mdss' => [],
      'rgws' => [],
    }

    ansible.extra_vars = {
      'common_single_host_mode' => true,
      'fsid'                    => '4a158d27-f750-41d5-9e7f-26ce4c9d2d45',
      'monitor_secret'          => 'AQAFx3RTAAAAABAAruXdSr8PTHAiTRgsyQMgPQ==',
      'cluster_network'         => '200.200.200.0/24',
      'public_network'          => '200.200.200.0/24',
      'monitor_interface'       => 'eth1',
      'devices'                 => ['/dev/sdb', '/dev/sdc'],
    }

    ansible.playbook = 'provision/provision.yml'
  end

end
