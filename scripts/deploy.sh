#!/bin/bash
ENDPOINT=${1:-8000}

PROJ_DIR=$(pwd)
CONFIG_DIR=${PROJ_DIR}/configs
TEMPLATE_DIR=${CONFIG_DIR}/server_skill_template
SKILL_DIR=${CONFIG_DIR}/server_skill
CONFIG_FILE=${CONFIG_DIR}/.serverconf.json
CONFIG_FILE_TEMPLATE=${CONFIG_DIR}/.serverconf.json.template

deploy_server() {
    bash -c "exec -a localserver make > /dev/null &"
}

get_skill_id() {
    cd ${SKILL_DIR}
    local skill_id=$(grep -hr "\"skill_id\"\:" | sed "s|\"skill_id\"\: \"\(.*\)\"\,|\1|g")
    cd ${PROJ_DIR}
    echo "${skill_id}"
}

modify_server_config() {
    local skill_id=$1 
    sed -i "s|\"alexaAppID\"\: \".*\"|\"alexaAppID\"\: \"$1\"|g" ${CONFIG_FILE}
}

build_server_config() {
    cp ${CONFIG_FILE_TEMPLATE} ${CONFIG_FILE}
}

deploy_ngrok() {
    ngrok http ${ENDPOINT} -log=stdout > /dev/null &
}

get_ngrok_url() {
    local address=$(curl --silent --show-error http://127.0.0.1:4040/api/tunnels | sed -nE 's/.*public_url":"https:..([^"]*).*/\1/p')
    echo "${address}"
}

set_ngrok_address_to_skill() {
    cd ${SKILL_DIR}
    grep -rl "\"uri\"\:" | xargs sed -i "s|\"uri\"\: \".*\"|\"uri\"\: \"https:\/\/$1\/alexa\"|g"
    cd ${PROJ_DIR}
}

build_skill() {
    if [[ ! -d "${SKILL_DIR}" ]]; then
        echo "no skill, creating."
        cd ${CONFIG_DIR}
        ask new -n "server_skill"
        rm -rf ${SKILL_DIR}/lambda
        cd ${PROJ_DIR}
    fi
    cp -a ${TEMPLATE_DIR}/. ${SKILL_DIR}/

}

deploy_skill() {
    echo "deploying skill"
    cd ${SKILL_DIR}
    ask deploy
    cd ${PROJ_DIR}
}

deploy_flow() {
    echo "Deploying the server"
    deploy_ngrok
    sleep 2s
    address=$(get_ngrok_url)
    echo "Ngrok address is ${address}"
    build_skill
    set_ngrok_address_to_skill ${address}
    deploy_skill
    skill_id=$(get_skill_id)
    echo "skill id is ${skill_id}"
    build_server_config
    modify_server_config ${skill_id}
    deploy_server
}


deploy_flow