#!/usr/bin/bash


es_wls=`ps -ef | grep weblogic | grep weblogic.home | grep -v grep | head -1`


if [ -z "$es_wls" ];then
    echo "not exists weblogic "
else
    wls_home=`echo $es_wls | awk '{ for(i=1; i<=NF; i++){ if ($i ~ /-Dweblogic.home/) print $i}}' | cut -d'=' -f2`
    echo ${wls_home}
fi



ps -ef | grep weblogic | grep AdminServer | awk '{ for(i=1; i<=NF; i++){ if ($i ~ /-Dweblogic.home/) print $i}}' | cut -d'=' -f2


/wls/wls11/utils/bsu/bsu.sh -prod_dir=/wls/wls11/wlserver_10.3 -status=applied -verbose -view

sudo su - root -c "cd /wls/wls11/utils/bsu;./bsu.sh -prod_dir=/wls/wls11/wlserver_10.3 -status=applied -verbose -view"

##WebLogic Server 10.3.6.0.200414 PSU Patch for BUG30857748 Sat Feb 29 22:41:51 PST 2020


getWls(){
    es_wls=`ps -ef | grep weblogic | grep weblogic.home | grep -v grep | head -1`
    echo ${es_wls}
}

getAdminWls(){
    admin_wls=`ps -ef | grep weblogic | grep AdminServer | grep -v grep | head -1`
    echo ${admin_wls}
}

# input wls base Dir  weblogic_home
getBsuPath(){
    ##base_dir=weblogic.home
    base_dir=$1
    lenth=`echo ${base_dir} | awk -F'/' '{print NF}'`
    for ((i=lenth; i>0;i--));
    do
        f_dir=`echo ${base_dir} | cut -d'/' -f1-${i}`
        rs=`sudo find ${f_dir} -name bsu.sh -type f 2>/dev/null`
        if [ `echo ${#rs}` -gt 3 ];then
            echo ${rs}
            break
        fi
    done
}


# input adminDomain Log Path
getLogPatch(){
    # find AdminDomain
    base_dir=$1
    cd $base_dir
    patch=""
    for f in $(find . -name "*.log*" -type f -size -500M  -printf "%T@ %Tc %p\n" | sort -n | awk '{print $NF}')
    do
        ex_pa=`cat $f | grep "Patch" | tail -1`
        if [ `echo ${#ex_pa}` -gt 5 ];then
            patch=`echo ${ex_pa}`
            
        fi
    done
    echo $patch
}

pch="$(getLogPatch)"

if [ `echo ${#pch}` -gt 5 ];then
    echo ${patch} | awk -F' ' '{print $3}'
fi