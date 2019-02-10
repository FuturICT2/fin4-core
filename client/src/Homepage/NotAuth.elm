module Homepage.NotAuth exposing (render)

import Common.Styles
    exposing
        ( padding
        , paddingBottom
        , paddingLeft
        , paddingTop
        , textCenter
        , toMdlCss
        )
import Homepage.Model exposing (Model)
import Homepage.Msg exposing (Msg(..))
import Html exposing (..)
import Html.Attributes exposing (style)
import Main.Context exposing (Context)
import Main.Routing exposing (loginPath, signupPath)
import Material.Button as Button
import Material.List as Lists


render : Context -> Model -> Html Msg
render ctx model =
    div [ paddingTop 30 ]
        [ Lists.ul [ toMdlCss (paddingLeft 25) ]
            [ Lists.li []
                [ Lists.content []
                    [ Lists.icon "people" []
                    , text "Engage with people"
                    ]
                ]
            , Lists.li []
                [ Lists.content []
                    [ Lists.icon "group_work" []
                    , text "Create value by sharing"
                    ]
                ]
            , Lists.li []
                [ Lists.content []
                    [ Lists.icon "money" []
                    , text "Monetize with cryptocurrency"
                    ]
                ]
            ]
        , div [ paddingLeft 96 ]
            [ Button.render Mdl
                [ 0 ]
                model.mdl
                [ Button.raised
                , Button.ripple
                , Button.colored
                , Button.link <| loginPath
                ]
                [ text "login" ]
            , div [ orStyle ] [ text " or  " ]
            , Button.render Mdl
                [ 1 ]
                model.mdl
                [ Button.raised
                , Button.ripple
                , Button.colored
                , Button.link <| signupPath
                ]
                [ text "join us" ]
            ]
        ]


orStyle : Attribute a
orStyle =
    style [ ( "margin", "0 15px" ), ( "display", "inline-block" ) ]
