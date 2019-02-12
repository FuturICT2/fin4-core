port module Asset.Ports exposing (ImagePortData, assetBlockImageRead, assetBlockImageSelected)


type alias ImagePortData =
    { contents : String
    , filename : String
    }


port assetBlockImageSelected : String -> Cmd msg


port assetBlockImageRead : (ImagePortData -> msg) -> Sub msg
