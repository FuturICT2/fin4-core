port module Actions.Ports exposing (ImagePortData, fileContentRead, fileSelected)


type alias ImagePortData =
    { contents : String
    , filename : String
    , tokenId : Int
    }


port fileSelected : String -> Cmd msg


port fileContentRead : (ImagePortData -> msg) -> Sub msg
