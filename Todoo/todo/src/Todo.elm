module Todo exposing (main)

import Browser
import Css exposing (..)
import Html.Styled exposing (..)
import Html.Styled.Attributes exposing (..)
import Html.Styled.Events exposing (onClick, onInput)
import VirtualDom


type alias User =
    { task : String
    , details : String
    , loggedIn : Bool
    }


initialModel : User
initialModel =
    { task = ""
    , details = ""
    , loggedIn = False
    }


view : User -> Html Msg
view user =
    div []
        [ h1 [ css [ paddingLeft (cm 3) ] ] [ text "Todo List" ]
        , styledForm []
            [ div []
                [ h3 [css [ textAlign (center) ] ] [ text "Task" ] 
                , styledInput
                    [ id "task"
                    , type_ "text"
                    , onInput SaveTask
                    ]
                    []
                ]
            , div []
                [ h3 [ css [ textAlign (center) ] ] [ text "Details" ]
                , styledDetails
                    [ id "details"
                    , type_ "details"
                    , onInput SaveDetails
                    ]
                    []
                ]
            , div []
                [ styledButton
                    [ type_ "submit"
                    , onClick Signup
                    ]
                    [ text "Add Todo" ]
                ]
            ]
        ]


styledForm : List (Attribute msg) -> List (Html msg) -> Html msg
styledForm =
    styled Html.Styled.form
        [ borderRadius (px 5)
        , backgroundColor (hex "#f2f2f2")
        , padding (px 20)
        , Css.width (px 300)
        ]


styledInput : List (Attribute msg) -> List (Html msg) -> Html msg
styledInput =
    styled Html.Styled.input
        [ display block
        , Css.width (px 260)
        , padding2 (px 12) (px 20)
        , margin2 (px 8) (px 0)
        , border (px 0)
        , borderRadius (px 4)
        ]

styledDetails : List (Attribute msg) -> List (Html msg) -> Html msg
styledDetails =
    styled Html.Styled.input
        [ display block
        , Css.width (px 260)
        , padding2 (px 50) (px 20)
        , margin2 (px 8) (px 0)
        , border (px 0)
        , borderRadius (px 4)
        ]


styledButton : List (Attribute msg) -> List (Html msg) -> Html msg
styledButton =
    styled Html.Styled.button
        [ Css.width (px 300)
        , backgroundColor (hex "#397cd5")
        , color (hex "#fff")
        , padding2 (px 14) (px 20)
        , marginTop (px 10)
        , border (px 0)
        , borderRadius (px 4)
        , fontSize (px 16)

        ]


type Msg
    = SaveTask String
    | SaveDetails String
    | Signup


update : Msg -> User -> User
update message user =
    case message of
        SaveTask task ->
            { user | task = task }

        SaveDetails details ->
            { user | details = details }

        Signup ->
            { user | loggedIn = True }


main : Program () User Msg
main =
    Browser.sandbox
        { init = initialModel
        , view = view >> toUnstyled
        , update = update
        }