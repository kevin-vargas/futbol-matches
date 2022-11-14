
export const verifySession = () => {
    if( !localStorage.getItem("token") ){
        window.location="/"
    }
}
