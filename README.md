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
    
    
    [root@test-display ~]# iostat
    Linux 3.10.0-327.el7.x86_64 (test-display) 	03/09/16 	_x86_64_	(32 CPU)
    
    avg-cpu:  %user   %nice %system %iowait  %steal   %idle
               1.68    0.00    3.48    0.06    0.00   94.78

    Device:            tps   Blk_read/s   Blk_wrtn/s   Blk_read   Blk_wrtn
    sda               0.00         0.00         0.00          0        171
    loop0             0.00         0.00         0.00       1433         16
    dm-0              0.00         0.00         0.00       1433         16
    dm-2              0.00         0.05         0.00      72947         32

    

# build
    go build -o cgroupfs github.com/adazhou16/cgroupfs/cli
    
    
    #test
    #test2
