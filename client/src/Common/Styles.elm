module Common.Styles exposing
    ( background
    , color
    , flexRow
    , float
    , margin
    , marginBottom
    , marginLeft
    , marginRight
    , marginTop
    , padding
    , paddingBottom
    , paddingLeft
    , paddingRight
    , paddingTop
    , pointer
    , textCenter
    , textJustify
    , textLeft
    , textRight
    , toMdlCss
    , toMdlCssList
    , width
    )

import Html exposing (..)
import Html.Attributes exposing (attribute, style)
import Material.Options


color : String -> Attribute a
color c =
    style [ ( "color", c ) ]


float : String -> Attribute a
float f =
    style [ ( "float", f ) ]


width : String -> Attribute a
width w =
    style [ ( "width", w ) ]


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


textJustify : Attribute a
textJustify =
    style [ ( "text-align", "justify" ) ]


padding : Int -> Attribute a
padding p =
    style [ ( "padding", toString p ++ "px" ) ]


paddingTop : Int -> Attribute a
paddingTop p =
    style [ ( "padding-top", toString p ++ "px" ) ]


paddingBottom : Int -> Attribute a
paddingBottom p =
    style [ ( "padding-bottom", toString p ++ "px" ) ]


paddingLeft : Int -> Attribute a
paddingLeft p =
    style [ ( "padding-left", toString p ++ "px" ) ]


paddingRight : Int -> Attribute a
paddingRight p =
    style [ ( "padding-right", toString p ++ "px" ) ]


margin : Int -> Attribute a
margin m =
    style [ ( "margin", toString m ++ "px" ) ]


marginTop : Int -> Attribute a
marginTop p =
    style [ ( "margin-top", toString p ++ "px" ) ]


marginBottom : Int -> Attribute a
marginBottom p =
    style [ ( "margin-bottom", toString p ++ "px" ) ]


marginLeft : Int -> Attribute a
marginLeft p =
    style [ ( "margin-left", toString p ++ "px" ) ]


marginRight : Int -> Attribute a
marginRight p =
    style [ ( "margin-right", toString p ++ "px" ) ]


flexRow : String -> String -> Attribute a
flexRow wrap justify =
    style
        [ ( "display", "flex" )
        , ( "flex-direction", "row" )
        , ( "flex-wrap", wrap )
        , ( "justify-content", justify )
        ]


background : String -> Attribute a
background b =
    style [ ( "background", b ) ]



{- converts normal elm styles to MDL css -}


toMdlCss : Attribute a -> Material.Options.Property c a
toMdlCss a =
    Material.Options.attribute a


toMdlCssList : List (Attribute a) -> List (Material.Options.Property c a)
toMdlCssList l =
    case l of
        h :: t ->
            toMdlCss h :: toMdlCssList t

        _ ->
            []
