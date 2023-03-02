const   AddHelpers = new LeaveAdd_Helpers(),
        AddAPI = new LeaveAdd_API()

window.addEventListener('DOMContentLoaded', () => {
    const myRIF = [ 'name', 'maxLimit', 'entitlementCalc', 'gender'],
        myForm = 'addLeaveForm'

    const myModal = document.querySelector('#openAddLeave')
    myModal.addEventListener('click', () => {
        Common.clearForm(myForm, myRIF)
    })

    const mySubmit = document.querySelector('#addLeaveSubmit')
    mySubmit.addEventListener('click', () =>{
        const error = Common.validateRequiredFields(myRIF)
        if (error == '0'){
            myData = Common.getForm(myForm)
            AddAPI.addLeave(myData).then(resp => {
                console.log(resp)
            })
        }
    })
    
    const showCard = true,
          hideCard = false,
          myCards = ['name', 'leaveConfig'];
    
    myCards.forEach(card => {
        document.querySelector(`#${card}Hide`).addEventListener('click', () => { Common.displayCardByID(card, hideCard) })
        document.querySelector(`#${card}Show`).addEventListener('click', () => { Common.displayCardByID(card, showCard) })
    })

    const myWarningMessage = document.querySelector('#hideWarningMessage')
    myWarningMessage.addEventListener('click', () => {
        Common.hideDivByID('warningMessageDiv')
    })
})
