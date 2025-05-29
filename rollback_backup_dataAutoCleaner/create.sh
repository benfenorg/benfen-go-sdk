#!/bin/bash

pwd

for ((i=0; i<25; i++))
do
    name="epoch_$i"
    # 在此处执行循环体的操作
    mkdir "./test_data/$name"
    echo $i
done
