curl -d 'client_id=istio' -d 'username=fexsor' -d 'password=fexsor' -d 'grant_type=password' 'http://keycloak:8080/auth/realms/demo/protocol/openid-connect/token' | python -m json.tool
