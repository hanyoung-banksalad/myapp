syntax = "proto3";

package v1.currency;

import "google/api/annotations.proto";
import "google/protobuf/wrappers.proto";

option go_package = "github.com/hanyoung-banksalad/myapp/idl/gen/myapp";
option java_outer_classname = "CurrencyProto";
option java_package = "com.github.banksalad.idl.apis.v1.currency";

service Currency {
  rpc HealthCheck(HealthCheckRequest) returns (HealthCheckResponse) {
    option (google.api.http) = {
      get: "/health"
    };
  }
  rpc GetExchangeRate(GetExchangeRateRequest) returns (GetExchangeRateResponse) {
    option (google.api.http) = {
      get: "/v1/currency/exchange-rate"
    };
  }
}

message HealthCheckRequest {}

message HealthCheckResponse {}

message GetExchangeRateRequest {
  // ISO 4217 Currency Code OR Non-ISO format string(ISO 4217 English name)
  string from_currency = 1;
  // NOTE: 현재는 KRW 로만 가능.
  // TODO 2020.06.30 이후로 supported currencies 추가 예정 (https://github.com/Rainist/idl/pull/623#issuecomment-640306480)
  string to_currency = 2;
  // 환율을 가져올 기준 날짜(int64 unix timestamp). 빈 값인 경우 서버의 현재 시간을 기준 환율을 반환합니다.
  google.protobuf.Int64Value base_date_ms = 3;
}

message GetExchangeRateResponse {
  // NOTE: 현재는 KRW 로만 가능.
  // TODO 2020.06.30 이후로 supported currencies 추가 예정 (https://github.com/Rainist/idl/pull/623#issuecomment-640306480)
  int64 exchange_rate_2f = 1;
}

