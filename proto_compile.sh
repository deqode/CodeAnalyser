
cd ..
protoc -I.  --go_out=.  code-analyser/protos/*/*/*.proto
protoc -I.  --go_out=.  code-analyser/protos/*/*.proto
cd code-analyser


