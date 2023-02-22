namespace go hello

struct pingRequest {
    1: string ping;
}

struct pingResponse {
    1: string pong;
}

service HelloService {
    pingResponse echo(1: pingRequest req)
}

// Model
struct ModelHello {
    1: i32 id (go.tag="gorm:\"primaryKey;column:id;index:idx_member,priority:1\"")
    2: i32 deleteAt (go.tag="gorm:\"primaryKey;column:delete_at;index:idx_member,priority:1\" json:\"id,omitempty\"")
}
