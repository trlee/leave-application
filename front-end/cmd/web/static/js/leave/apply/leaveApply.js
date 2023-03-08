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

    const uploadButton = document.getElementById("uploadButton");
    uploadButton.addEventListener('click', (event) => {
        event.preventDefault()
        let fileInput = document.getElementById("attachments")
        let files = fileInput.files;
        const formData = new FormData();

        for (let i = 0; i < files.length; i++) {
            let file = files[i];
            formData.append("attachments[]", file);
        }
        API.submitAttachment(formData).then(resp => {
            console.log(resp.data)
            if (!resp.error) {
                Helpers.populateInputHidden("filenamesTobeInsert", resp.data)
                console.log(resp.data)
                alert("Successfully uploaded")
            } else {
                alert("Fail to upload")
            }
        })
    })
    
    const showCard = true,
    hideCard = false,
    myCards = ['employeeInfo', 'leaveInfo', 'attachment'];
    
    myCards.forEach(card => {
        document.querySelector(`#${card}Hide`).addEventListener('click', () => { Common.displayCardByID(card, hideCard) })
        document.querySelector(`#${card}Show`).addEventListener('click', () => { Common.displayCardByID(card, showCard) })
    })
    
    const myWarningMessage = document.querySelector('#hideWarningMessage')
    myWarningMessage.addEventListener('click', () => {
        Common.hideDivByID('warningMessageDiv')
    })
})