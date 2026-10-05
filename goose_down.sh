#! /bin/bash

cd sql/schema
goose postgres postgres://myuser:mypassword@192.168.1.210:5432/chirpy down
cd -
