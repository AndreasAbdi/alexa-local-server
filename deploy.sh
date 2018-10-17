#!/bin/bash
ENDPOINT=${1:-8000}

deploy_server() {
    bash -c "exec -a localserver make > /dev/null &"
}

deploy_ngrok() {
    ngrok http ${ENDPOINT} -log=stdout > /dev/null &
}

get_ngrok_url() {
    local address=$(curl --silent --show-error http://127.0.0.1:4040/api/tunnels | sed -nE 's/.*public_url":"https:..([^"]*).*/\1/p')
    echo "${address}"
}

set_ngrok_address_to_skill() {
    grep -rl --exclude-dir=server_skill_template --exclude-dir=.git --exclude-dir=vendor '"uri":' | xargs sed -i "s|\"uri\"\: \".*\"|\"uri\"\: \"https:\/\/$1\"|g"
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
}

deploy_server
deploy_ngrok
sleep 2s
address=$(get_ngrok_url)
echo "Ngrok address is ${address}"
build_skill
set_ngrok_address_to_skill ${address}
deploy_skill