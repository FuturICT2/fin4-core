module Actions.Sub exposing (subscriptions)

import Actions.Msg exposing (Msg(..))
import Actions.Ports exposing (fileContentRead)
import Main.Context exposing (Context)
import Time


subscriptions : Context -> Sub Msg
subscriptions ctx =
    Sub.batch
        [ Time.every 500 (always TickerTimout)
        , fileContentRead ImageRead
        ]



-- Sub.none
-- Time.every 500 (always TickerTimout)
