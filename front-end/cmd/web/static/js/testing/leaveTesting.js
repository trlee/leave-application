const API = new Leave_API()

window.addEventListener('DOMContentLoaded', (event) => {
    const myBtn = document.querySelector('#submit_button'),
        myEmail = document.querySelector('#email_input'),
        myDuration = document.querySelector('#duration_input'),
        myReason = document.querySelector('#reason_input'),
        myHalfDay = document.querySelector('#half_day_input')

    myBtn.addEventListener('click', ()=> {
        API.verifyLeave(myEmail.value, myDuration.value, myReason.value).then(resp => {
            console.log('resp')
            console.log(resp)

            if (!resp.error){
                console.log('Successfully logged')
            }else{
                console.log('Something went wrong')
            }
        })
    })
    
})