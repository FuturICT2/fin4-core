module Token.Sub exposing (subscriptions)

import Main.Context exposing (Context)
import Time
import Token.Msg exposing (Msg(..))
import Token.Ports exposing (fileContentRead)


subscriptions : Context -> Sub Msg
subscriptions ctx =
    fileContentRead ImageRead



-- Sub.none
-- Time.every 500 (always TickerTimout)
