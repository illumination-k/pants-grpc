```bash
printf '[GLOBAL]\npants_version = "2.11.0"\n' > pants.toml
curl -L -O https://static.pantsbuild.org/setup/pants && chmod +x ./pants

go mod init github.com/illumination-k/pants-grpc
mkdir -p 3rdparty/python
cat <<EOF > 3rdparty/python/requirements.txt
grpcio>=1.46.0
protobuf~=3.19.0
EOF

mkdir -p client server protos
touch client/main.go server/main.py protos/user.proto

./pants tailor
./pants generate-lockfiles

./pants run client:bin
./pants run server/main.py
```