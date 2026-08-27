# core

```txt
service UserService {
  // 基础操作
  rpc CreateUser(CreateUserRequest) returns (CreateUserResponse);
  rpc UpdateUser(UpdateUserRequest) returns (UpdateUserResponse);
  rpc GetUser(GetUserRequest) returns (GetUserResponse);
  rpc ListUsers(ListUsersRequest) returns (ListUsersResponse);
  rpc DeleteUser(DeleteUserRequest) returns (DeleteUserResponse);

  // 扩展操作
  rpc BatchDeleteUsers(BatchDeleteUsersRequest) returns (BatchDeleteUsersResponse);
  rpc EnableUser(EnableUserRequest) returns (EnableUserResponse);
  rpc DisableUser(DisableUserRequest) returns (DisableUserResponse);
  rpc ResetUserPassword(ResetUserPasswordRequest) returns (ResetUserPasswordResponse);
}

service UserService {
  // 基础操作
  rpc Create(CreateUserRequest) returns (CreateUserResponse);
  rpc Update(UpdateUserRequest) returns (UpdateUserResponse);
  rpc Get(GetUserRequest) returns (GetUserResponse);
  rpc List(ListUsersRequest) returns (ListUsersResponse);
  rpc Delete(DeleteUserRequest) returns (DeleteUserResponse);

  // 业务操作
  rpc Enable(EnableUserRequest) returns (EnableUserResponse);
  rpc Disable(DisableUserRequest) returns (DisableUserResponse);
  rpc ResetPassword(ResetUserPasswordRequest) returns (ResetUserPasswordResponse);
  rpc AssignRoles(AssignUserRolesRequest) returns (AssignUserRolesResponse);
}


platform-base/
rpc/proto/platform_base.proto

platform-operator/
rpc/proto/platform_operator.proto

platform-game/
rpc/proto/platform_game.proto

platform-message/
rpc/proto/platform_message.proto

integration/
rpc/proto/integration.proto

file/
rpc/proto/file.proto

node-dispatch/
rpc/proto/node_dispatch.proto


platform-base/
├── api/
│   └── desc/
│       └── main.api
│
└── rpc/
    └── proto/
        ├── platform_base.proto
        └── types/
            ├── common.proto
            ├── user.proto
            ├── role.proto
            └── ...
```