module Common.Styles exposing (..)

import Html exposing (..)
import Html.Attributes exposing (style, attribute)
import Material.Options


pointer : Attribute a
pointer =
    style [ ( "cursor", "pointer" ) ]


textLeft : Attribute a
textLeft =
    style [ ( "text-align", "left" ) ]


textRight : Attribute a
textRight =
    style [ ( "text-align", "right" ) ]


textCenter : Attribute a
textCenter =
    style [ ( "text-align", "center" ) ]


padding : Int -> Attribute a
padding p =
    style [ ( "padding", (toString p) ++ "px" ) ]


margin : Int -> Attribute a
margin m =
    style [ ( "margin", (toString m) ++ "px" ) ]


flexRow wrap justify =
    style
        [ ( "display", "flex" )
        , ( "flex-direction", "row" )
        , ( "flex-wrap", wrap )
        , ( "justify-content", justify )
        ]



{- converts normal elm styles to MDL css -}


toMdlCss : Attribute a -> Material.Options.Property c a
toMdlCss a =
    Material.Options.attribute a
