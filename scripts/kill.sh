#!/bin/sh

ps aux | grep localserver | awk '{print $2}' | xargs kill

ps aux | grep ngrok | awk '{print $2}' | xargs kill

ps aux | grep alexa_local_server | awk '{print $2}' | xargs kill
