#!/bin/bash
# 用法: ./new_practice.sh <编号> <主题名>
# 示例: ./new_practice.sh 1 variable
#       ./new_practice.sh 10 slice

set -e

if [ $# -ne 2 ]; then
    echo "用法: $0 <编号> <主题名>"
    echo "示例: $0 1 variable"
    echo "      $0 10 slice"
    exit 1
fi

NUM=$1
TOPIC=$2

# 将主题名首字母大写，用于测试函数名
FUNC_NAME=$(echo "$TOPIC" | awk '{print toupper(substr($0,1,1)) substr($0,2)}')

FILE_NAME="${NUM}_${TOPIC}_test.go"
FILE_PATH="$(dirname "$0")/${FILE_NAME}"

if [ -f "$FILE_PATH" ]; then
    echo "❌ 文件已存在: $FILE_NAME"
    exit 1
fi

cat > "$FILE_PATH" <<EOF
package practice

import (
	"fmt"
	"testing"
)

func Test${FUNC_NAME}(t *testing.T) {
	// TODO: 在这里写你的练习代码
	fmt.Println("practice: ${TOPIC}")
}
EOF

echo "✅ 已创建: $FILE_NAME"
echo "   函数名: Test${FUNC_NAME}"
echo "   运行命令: go test -v -run Test${FUNC_NAME}"
