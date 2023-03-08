namespace go oss

struct PreSignedPutObjectUrlReq {
    1: string bucket_name;
    2: string object_name;
    3: i32 expiry; // 有效期：单位为 s ，默认为 7 days
}

struct PreSignedPutObjectUrlResponse {
    1: string pre_signed_url;
}

service OssService {
    // 生成上传的 url 地址
    PreSignedPutObjectUrlResponse PreSignedPutObjectUrl(1: PreSignedPutObjectUrlReq req)
}
