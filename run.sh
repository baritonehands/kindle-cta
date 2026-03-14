#!/bin/sh

export CTA_TRAIN_TRACKER_API_KEY=<paste>
export CTA_BUS_TRACKER_API_KEY=<paste>

/usr/bin/lipc-set-prop -- com.lab126.powerd preventScreenSaver 1

./myapp > /var/log/kindle-cta.log

/usr/bin/lipc-set-prop -- com.lab126.powerd preventScreenSaver 0
