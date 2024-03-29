namespace go oss

// OssType oss 存储类型
enum OssPlatformType {
    UNKNOWN = 0,
    MINIO = 1,
    ALI_YUN = 2,
}

struct PreSignedPutObjectUrlReq {
    1: string bucket_name;
    2: string object_name;
    3: i32 expiry; // 有效期：单位为 s ，默认为 7 days
    4: OssPlatformType type; // 存储平台类型
}

struct PreSignedPutObjectUrlResp {
    1: string pre_signed_url;
    2: string resource_url; // 上传后的资源路径
}

struct PutObjectReq {
    1: string bucket_name;
    2: string object_name;
    3: binary file;
    4: OssPlatformType type; // 存储平台类型
}

struct PutObjectResp {
    1: string url;
}

service OssService {
    // 生成上传的 url 地址
    PreSignedPutObjectUrlResp PreSignedPutObjectUrl(1: PreSignedPutObjectUrlReq req);
    // 通过二进制上传文件
    PutObjectResp PutObject(1: PutObjectReq req);
}
