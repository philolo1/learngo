#!/bin/bash
aws lambda add-permission --function-name goTranslator \
    --action lambda:InvokeFunctionUrl \
    --principal "*" \
    --function-url-auth-type "NONE" \
    --statement-id url

aws lambda create-function-url-config \
    --function-name goTranslator \
    --auth-type NONE
