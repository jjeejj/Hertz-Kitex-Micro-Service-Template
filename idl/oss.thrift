namespace go oss

// OssType oss 存储类型
enum OssType {
    MINIO = 1,
    ALI_YUN = 2,
}

struct PreSignedPutObjectUrlReq {
    1: string bucket_name;
    2: string object_name;
    3: i32 expiry; // 有效期：单位为 s ，默认为 7 days
    4: OssType type; // 存储类型
}

struct PreSignedPutObjectUrlResponse {
    1: string pre_signed_url;
}

service OssService {
    // 生成上传的 url 地址
    PreSignedPutObjectUrlResponse PreSignedPutObjectUrl(1: PreSignedPutObjectUrlReq req)
}
