#!/bin/bash
# root user execute
#
ps -eo pid,user,%cpu | sort -nk4 | awk '{if($3>50)print}' | while read pid user cpu; do
    if [ $pid -gt 0 ]; then
        echo -e "\n\n\n\t\t\t\tProcess: ${pid} User: ${user} CPU(%): ${cpu}% \n\n\n"

        echo "SubProcess INFO..............."
        ps -p ${pid} -o pid,ppid,fgroup,ni,lstart,etime,cmd
        echo -e "\nFD INFO..................\n"
        ls -l /proc/${pid}/fd/
        echo -e "\nNetwork Connect ..............\n"
        netstat -anp | grep ${pid}

        echo -e "\n Crontab Check Last Two Days Write Cron.............\n"
        find /var/spool/cron -mtime -2 -type f
    else
        echo "[INFO] Not Found CPU>50% Process...."
    fi
done
