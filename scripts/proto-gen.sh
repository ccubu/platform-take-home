echo "Generating proto code"


function check_command() {
  if ! command -v "$1" &>/dev/null; then
    echo "Error: '$1' not installed. Please install it and try again."
    exit 1
  fi
}

# Check for needed tools
check_command buf
check_command protoc
check_command protoc-gen-go
check_command protoc-gen-go-grpc
check_command protoc-gen-go-grpc-mock
check_command protoc-gen-grpc-gateway


cd proto
proto_dirs=$(find . -path -prune -o -name '*.proto' -print0 | xargs -0 -n1 dirname | sort | uniq)
for dir in $proto_dirs; do
  for file in $(find "${dir}" -maxdepth 1 -name '*.proto'); do
    if grep go_package $file &>/dev/null; then
      buf generate --template buf.gen.yaml $file
    fi
  done
done

# Generate external protocol buffers
# echo "Generating cosmwasm protos"
# buf generate buf.build/cosmwasm/wasmd

rm -r ../api/types
mv github.com/skip-mev/platform-take-home/api/types ../api
rm -r github.com

# move proto files to the right places

echo "Proto code generation completed."