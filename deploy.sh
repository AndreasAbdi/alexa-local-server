#!/bin/bash
ENDPOINT=${1:-8000}

deploy_server() {
    bash -c "exec -a localserver make > /dev/null &"
}

get_skill_id() {
    local skill_id=$(grep -hr --exclude-dir=server_skill_template --exclude-dir=.git --exclude-dir=vendor "\"skill_id\"\:" | sed "s|\"skill_id\"\: \"\(.*\)\"\,|\1|g")
    echo "${skill_id}"
}

modify_server_config() {
    local skill_id=$1 
    sed -i "s|\"alexaAppID\"\: \".*\"|\"alexaAppID\"\: \"$1\"|g" ./.serverconf.json
}

build_server_config() {
    cp ./.serverconf.json.template ./.serverconf.json
}

deploy_ngrok() {
    ngrok http ${ENDPOINT} -log=stdout > /dev/null &
}

get_ngrok_url() {
    local address=$(curl --silent --show-error http://127.0.0.1:4040/api/tunnels | sed -nE 's/.*public_url":"https:..([^"]*).*/\1/p')
    echo "${address}"
}

set_ngrok_address_to_skill() {
    grep -rl --exclude-dir=server_skill_template --exclude-dir=.git --exclude-dir=vendor "\"uri\"\:" | xargs sed -i "s|\"uri\"\: \".*\"|\"uri\"\: \"https:\/\/$1\"|g"
}

build_skill() {
    if [[ ! -d "./server_skill" ]]; then
        echo "no skill, creating."
        ask new -n "server_skill"
        rm -rf ./server_skill/lambda
        cp -a ./server_skill_template/. ./server_skill/
    fi
}

deploy_skill() {
    cd ./server_skill
    ask deploy
    cd ../
    pwd
}

echo $(pwd)
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
