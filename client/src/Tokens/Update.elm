module Tokens.Update exposing (update)

import Debug
import Main.Context exposing (Context)
import Material
import Model.Tokens exposing (tokensDecoder)
import Tokens.Command exposing (loadTokensCmd)
import Tokens.Model exposing (Model)
import Tokens.Msg exposing (Msg(..))


update : Context -> Msg -> Model -> ( Model, Cmd Msg )
update ctx msg model =
    case msg of
        OnLoadTokensResponse resp ->
            case resp of
                Ok tokens ->
                    { model | tokens = Just tokens, error = Nothing } ! []

                Err error ->
                    Debug.log (toString error)
                        { model
                            | tokens = Nothing
                            , error = Just "Error loading tokens, please try again!"
                        }
                        ! []

        SelectTab tab ->
            { model | selectedTab = tab } ! []

        TickerTimout ->
            model ! [ loadTokensCmd ctx 1 ]

        Mdl msg_ ->
            Material.update Mdl msg_ model
