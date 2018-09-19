module Common.Error exposing (render, renderMaybeError, renderHttpError, renderWebDataError)

import RemoteData exposing (WebData, RemoteData)
import Html.Attributes exposing (..)
import Material.Spinner as Spinner
import Html exposing (..)
import Http exposing (Error, Error(..))


noError : Html a
noError =
    div [] []


render : String -> Html a
render error =
    case error of
        "" ->
            noError

        _ ->
            div [ errorStyle ]
                [ text error
                ]


renderMaybeError : Maybe String -> Html a
renderMaybeError maybeError =
    case maybeError of
        Just error ->
            render error

        Nothing ->
            noError


renderHttpError : Maybe Error -> Html a
renderHttpError error =
    case error of
        Just httpError ->
            case httpError of
                BadStatus errorObj ->
                    render errorObj.body

                _ ->
                    noError

        Nothing ->
            noError


renderWebDataError : WebData a -> Html b
renderWebDataError webdata =
    case webdata of
        RemoteData.Failure error ->
            renderHttpError <| Just error

        _ ->
            noError


errorStyle =
    style
        [ ( "color", "rgb(255,82,82)" )
        , ( "padding", "10px 0" )
        ]
