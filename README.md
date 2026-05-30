# ติดตั้ง xk6
go install go.k6.io/xk6/cmd/xk6@latest

# build k6 รวม extension เข้าไป
cd /Users/sarawutt.bur/k6-custom

xk6 build v1.7.1 \
  --with github.com/dkantikorn/k6-x-counter=./k6-x-counter \
  --output ./k6-custom

# รันด้วย binary ใหม่แทน k6 ปกติ
./k6-custom run script.js