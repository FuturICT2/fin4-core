module Main exposing (init, main)

import Main.Command exposing (commands)
import Main.Flags exposing (Flags)
import Main.Model exposing (Model, initModel)
import Main.Msg exposing (Msg(..))
import Main.Routing exposing (parseLocation)
import Main.Sub exposing (subscriptions)
import Main.Update exposing (update)
import Main.View exposing (view)
import Navigation exposing (Location)


init : Flags -> Location -> ( Model, Cmd Msg )
init flags location =
    let
        -- initialize model
        model_ =
            initModel flags <| parseLocation location

        -- trigger initial OnRouteChange
        ( model, cmds ) =
            update (OnRouteChange location) model_
    in
    ( model, Cmd.batch [ cmds, commands model.context ] )


main : Program Flags Model Msg
main =
    Navigation.programWithFlags OnRouteChange
        { init = init
        , view = view
        , update = update
        , subscriptions = subscriptions
        }
