#!/usr/bin/env bash

# Clusbreak statuses are implemented as:
# https://github.com/hortonworks/cloudbreak/blob/master/core-api/src/main/java/com/sequenceiq/cloudbreak/api/model/Status.java
# REQUESTED,
# CREATE_IN_PROGRESS,
# AVAILABLE,
# UPDATE_IN_PROGRESS,
# UPDATE_REQUESTED,
# UPDATE_FAILED,
# CREATE_FAILED,
# ENABLE_SECURITY_FAILED,
# PRE_DELETE_IN_PROGRESS,
# DELETE_IN_PROGRESS,
# DELETE_FAILED,
# DELETE_COMPLETED,
# STOPPED,
# STOP_REQUESTED,
# START_REQUESTED,
# STOP_IN_PROGRESS,
# START_IN_PROGRESS,
# START_FAILED,
# STOP_FAILED,
# WAIT_FOR_SYNC

function wait-cluster-status() {
    is_status="false"

    while [ $SECONDS -lt $1 ] && [ "${is_status}" == "false" ]
    do
    	sleep 30
	    is_status=$(cb cluster describe --name $2 | jq -r '."status" == "$3"')
    done

    if [[ "${is_status}" != "true" ]]; then
        echo $(cb cluster describe --name $2 | jq -r '."statusReason"')
    else
        echo "${is_status}"
    fi
}

function cluster-is-status() {
    is_status="false"

	is_status=$(cb cluster describe --name $1 | jq -r '."status" == "$2"')

    if [[ "${is_status}" != "true" ]]; then
        echo $(cb cluster describe --name $1 | jq -r '."statusReason"')
    else
        echo "${is_status}"
    fi
}

function wait-stack-cluster-status() {
    param1=$3
    export param1
    is_status=false
    has_failure=false
    while [ $SECONDS -lt $1 ] && [ $is_status == false ] && [ $has_failure == false ]
    do
	sleep 30
	stack_status=$(describe-cluster --name $2 | jq -r '."status"')
	cluster_status=$(describe-cluster --name $2 | jq -r ' ."cluster" | ."status"')

	is_stack_status=$(describe-cluster --name $2 | jq -r '."status"==env.param1')
	is_cluster_status=$(describe-cluster --name $2 | jq -r ' ."cluster" | ."status"==env.param1')

	is_status= is_cluster_status && is_stack_status
	is_status=true

    status_reason_stack=$(describe-cluster --name $2 | jq -r '."status"==env.param1')
	status_reason_cluster=$(describe-cluster --name $2 | jq -r ' ."cluster" | ."statusReason"')

    if [ $cluster_status = UPDATE_FAILED ] || [ $cluster_status = CREATE_FAILED ] || [ $stack_status = UPDATE_FAILED ] || [ $stack_status = CREATE_FAILED ]; then
      has_failure=true
    fi
    done
    echo $status_reason_stack >> /tmp/status_reason
	echo $status_reason_cluster >> /tmp/status_reason
    [ $is_status == true ]
}

function stack-cluster-is-status() {
    param1=$2
    export param1
    is_status=false
    is_status_stack=$(describe-cluster --name $1 | jq -r '."status"==env.param1')
	is_status_cluster=$(describe-cluster --name $1 | jq -r ' ."cluster" | ."status"==env.param1')
	is_status= is_status_stack && is_status_cluster
    [ $is_status == true ]
}

function node-count-are-equal() {
    NODES_COUNT_CLUSTER=$( describe-cluster --name $1 | jq ' ."instanceGroups" | .[] | select(."group"=="compute") | . "nodeCount" ')
    NODES_COUNT_STACK=$( describe-cluster --name $1 | jq ' ."cluster"  | . "hostGroups" | . [] | . "constraint" | select(."instanceGroupName" == "compute") | . "hostCount"')
    [ $NODES_COUNT_CLUSTER = $NODES_COUNT_STACK ]
}

function wait-cluster-delete() {
    param1=$2
    export param1
    is_status=false
    while [ $SECONDS -lt $1 ] && [ $is_status == false ]
    do
	sleep 30
    tmp=$(list-clusters | jq '.[] |  select(."Name" == env.param1) ' )
	if [$tmp == ""] ; then
	   is_status=true
	fi
    done
    echo $is_status
    [ $is_status == true ]
}

function remove-stuck-cluster() {
  if [[ $(cb cluster list | jq -r '.[].Name' | grep $1) ]]; then
    cb cluster delete --name $1 --wait
  fi
}

function remove-stuck-credential() {
  if [[ $(cb credential list | jq -r '.[].Name' | grep $1) ]]; then
    cb credential delete --name $1
  fi
}
