// Generated by protoc-gen-twirp-swift. DO NOT EDIT
{{- range .Imports}}
import {{.Path}};
{{- end}}

enum TwirpError:Error {
    case httpError(withStatus:Int)
    case decodeError(withData:String? = nil)
}

class {{.ClientName}} {
    let requestHandler:RequestHandler
    let baseUrl: URL

    init(requestHandler:RequestHandler = DefaultRequestHandler(), baseUrl: String){
        self.requestHandler = requestHandler
        self.baseUrl = URL(string: baseUrl)!
    }

    func decodeMessage<M:Message>(fromJson jsonData:Data?) throws -> M {
        guard let data = jsonData else {
            throw TwirpError.decodeError()
        }

        do {
            return try M(jsonUTF8Data: data)
        } catch is BinaryDecodingError {
            throw TwirpError.decodeError(withData: String(decoding: data, as: UTF8.self))
        }

    }

    func decodeMessage<M:Message>(fromProto protobufData:Data?) throws -> M {
        guard let data = protobufData else {
            throw TwirpError.decodeError()
        }
        do {
            return try M(serializedData: data)
        } catch is BinaryDecodingError {
            throw TwirpError.decodeError(withData: String(decoding: data, as: UTF8.self))
        }
    }

    func perform<M:Message>(_ request: Message, withPath endpointSuffix: String) async throws -> M {
        let url = baseUrl.appendingPathComponent(endpointSuffix)
        let req = requestHandler.createRequest(url:url, message:request)
        let data = try await requestHandler.performRequest(req)
        
        return try decodeMessage(fromProto: data)
    }

}

protocol RequestHandler {
    func createRequest(url:URL, message:Message) -> URLRequest
    func performRequest(_ req:URLRequest) async throws  -> Data
}


class DefaultRequestHandler: RequestHandler {
    var session = URLSession.shared

    func createRequest(url: URL, message: Message) -> URLRequest {
        var req = URLRequest.init(url:url);
        req.setValue("application/protobuf", forHTTPHeaderField: "Content-Type")
        req.httpMethod = "POST"

        let data = try! message.serializedData()
        req.httpBody = data
        return req
    }

    func performRequest(_ req: URLRequest) async throws -> Data {
        let (responseData, _) = try await session.data(
            for: req
        )
        return responseData
    }
}

{{range .Services}}
protocol {{.ClassName}} {
    {{range .Methods}}
	func {{.Name}}(_: {{.InputType}}) async throws -> {{.OutputType}}
    {{end}}
}
{{end}}

{{range .Services}}
extension {{$.ClientName}}: {{.ClassName}} {
    var {{.ServicePathPrefixName}}: String { return "/twirp/{{.Package}}.{{.Name}}/" }
    {{$s := .ServicePathPrefixName }}
    {{range .Methods}}
    func {{.Name}}(_ request: {{.InputType}}) async throws -> {{.OutputType}} {
        return try await perform(request, withPath: {{$s}}+"{{.Path}}")
    }
    {{end}}
}
{{end}}