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
    , user = Nothing
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
