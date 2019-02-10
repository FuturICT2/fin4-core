module Asset.Sub exposing (subscriptions)

import Asset.Msg exposing (Msg(..))
import Asset.Ports exposing (assetBlockImageRead)
import Main.Context exposing (Context)


subscriptions : Context -> Sub Msg
subscriptions ctx =
    Sub.batch
        [ assetBlockImageRead ImageRead
        ]
