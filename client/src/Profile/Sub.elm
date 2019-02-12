module Profile.Sub exposing (subscriptions)

import Main.Context exposing (Context)
import Profile.Msg exposing (Msg(..))
import Profile.Ports exposing (proileImageContentRead)


subscriptions : Context -> Sub Msg
subscriptions ctx =
    Sub.batch
        [ proileImageContentRead ImageRead
        ]
