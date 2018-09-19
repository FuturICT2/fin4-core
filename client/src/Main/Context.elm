module Main.Context exposing (Context, DeviceSize(..), initContext, isUserId)

import Main.Flags exposing (Flags)
import Main.Routing exposing (Route)
import Main.User exposing (User, noUserYet)


type DeviceSize
    = Mobile
    | Tablet
    | Desktop


type alias Context =
    { route : Route
    , flags : Flags
    , user : Maybe User
    , sessionDidLoad : Bool
    , window :
        { width : Int
        , height : Int
        , device : DeviceSize
        , isMobile : Bool
        }
    }


initContext : Flags -> Route -> Context
initContext flags route =
    { route = route
    , flags = flags
    , sessionDidLoad = False
    , user =
        Just
            { id = 1
            , email = "jouda@fin4.com"
            , address = "0xc98e86927d9752586da1081c8dd9a41450232deb"
            }
    , window =
        { width = 0
        , height = 0
        , device = Desktop
        , isMobile = False
        }
    }


isUserId : Context -> Int -> Bool
isUserId ctx id =
    case ctx.user of
        Just user ->
            id == user.id

        _ ->
            False
