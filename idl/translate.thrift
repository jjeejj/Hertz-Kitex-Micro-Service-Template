// 翻译服务
namespace go translate

// 翻译平台类型
// 阿里云的语言 code https://help.aliyun.com/document_detail/215387.html?spm=a2c4g.11186623.0.0.793b4f62itNjSI
enum TranslatePlatformType {
    UNKNOWN = 0,
    ALI_YUN = 1, // 阿里云
}

struct DetectLanguageReq {
    1: string SourceText;
    2: TranslatePlatformType type;
}

struct DetectLanguageResp {
    1: string DetectedLanguage;
}

struct TranslateReq {
    1: string SourceText;
    2: string SourceLanguage; // 支持自动检测，语言代码为auto
    3: string TargetLanguage;
    4: TranslatePlatformType type;
}

struct TranslateResp {
    1: string translateResult;
}

service TranslateService {
    // 识别指定文本语言
   DetectLanguageResp DetectLanguage(1: DetectLanguageReq req);
   // 翻译文本
   TranslateResp Translate(1: TranslateReq req);
}
