port module Profile.Ports exposing
    ( ImagePortData
    , profileImageSelected
    , proileImageContentRead
    )


type alias ImagePortData =
    { contents : String
    , filename : String
    }


port profileImageSelected : String -> Cmd msg


port proileImageContentRead : (ImagePortData -> msg) -> Sub msg
