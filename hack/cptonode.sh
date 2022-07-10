#!/bin/bash
file=whatever
kubectl run busybox --image=busybox:1.34 --command -- sleep 36000
kubectl cp ./($file) default/busybox:/($file)
//enter the node
kubectl cp default/busybox:/($file) ./($file)