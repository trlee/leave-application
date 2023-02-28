const   Common = new Main_Helpers(),
        Helpers = new LeaveApply_Helpers(),
        API = new LeaveApply_API()

window.addEventListener('DOMContentLoaded', () => {
    // Populate Function for Leave Application (TBD: Put everything here into a base layout for leave pages)
    const myRIF = [ 'email', 'leaveType', 'leaveDate', 'halfDay', 'applyReason'],
    myForm = 'applyLeaveForm'
    
    const myModal = document.querySelector('#openApplyLeave')
    myModal.addEventListener('click', () => {
        Common.clearForm(myForm, myRIF)
    })
    
    const mySubmit = document.querySelector('#applyLeaveSubmit')
    mySubmit.addEventListener('click', () =>{
        const error = Common.validateRequiredFields(myRIF)
        if (error == '0'){
            myData = Common.getForm(myForm)
            API.applyLeave(myData).then(resp => {
                console.log(resp)
            })
        }
    })
    
    const showCard = true,
    hideCard = false,
    myCards = ['employeeInfo', 'leaveInfo'];
    
    myCards.forEach(card => {
        document.querySelector(`#${card}Hide`).addEventListener('click', () => { Common.displayCardByID(card, hideCard) })
        document.querySelector(`#${card}Show`).addEventListener('click', () => { Common.displayCardByID(card, showCard) })
    })
    
    const myWarningMessage = document.querySelector('#hideWarningMessage')
    myWarningMessage.addEventListener('click', () => {
        Common.hideDivByID('warningMessageDiv')
    })
})