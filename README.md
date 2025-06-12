to send:
curl -X POST http://localhost:8000/otp \
  -H "Content-Type: application/json" \
  -d '{
    "phoneNumber": "+9134567890"
  }'


to verify:
curl -X POST http://localhost:8000/verifyOTP \
  -H "Content-Type: application/json" \
  -d '{
    "user": {
      "phoneNumber": "+9134567890"
    },
    "code": "123456"
  }'
