# cgroupfs

Like lxcfs https://github.com/lxc/lxcfs, cgroupfs provides an emulated /proc/meminfo, /proc/cpuinfo... for containers.

# Usage

    container_id=`docker create --device /dev/fuse --cap-add SYS_ADMIN -v /tmp/cgroupfs/meminfo:/proc/meminfo -m=15m ubuntu sleep 213133`

    ## in the second console tab
    go run cli/cli.go /tmp/cgroupfs /docker/$container_id

    ## go to the first tab
    docker start $container_id

    [root@test-display ~]# free -g
             total       used       free     shared    buffers     cached
    Mem:             8          0          7          0          0          0
    -/+ buffers/cache:          0          7 
    Swap:            0          0          0 
    
    

# build
    go build -o cgroupfs github.com/adazhou16/cgroupfs/cli
