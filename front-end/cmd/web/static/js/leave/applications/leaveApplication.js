const   ApplicationHelpers = new LeaveApplication_Helpers(),
        ApplicationAPI = new LeaveApplication_API(),
        // Set datatable parameters (JQuery Datatable)
         dt = 'leaveApplication',
         dtBody = 'leaveApplicationBody'

window.addEventListener('DOMContentLoaded', () => {
    // fetch all leave applications & update DOM    
    ApplicationAPI.getAllLeaveApplication().then(resp => { 
        console.log(resp)
        ApplicationHelpers.insertRows(dtBody, resp.data)
        ApplicationHelpers.triggerDT(dt, dtBody)
    })


    myForm = 'updateLeaveAppForm'

    // Submitting Cancel Leave + Reason
    const myUpdateCancel = document.querySelector('#cancelLeaveSubmit')
    myUpdateCancel.addEventListener('click', () => {
        console.log('In Cancel')
        myData = Common.getForm(myForm)
        ApplicationAPI.cancelLeaveApplication(myData).then(resp => {
            console.log(resp)
        })
    })

    // Submitting Approve Leave + Reason
    const myUpdateApprove = document.querySelector('#approveLeaveSubmit')
    myUpdateApprove.addEventListener('click', () => {
        console.log('In Approve')
        myData = Common.getForm(myForm)
        ApplicationAPI.approveLeaveApplication(myData).then(resp => {
            console.log(resp)
        })
    })
    
    // Submitting Reject Leave + Reason
    const myUpdateReject = document.querySelector('#rejectLeaveSubmit')
    myUpdateReject.addEventListener('click', () => {
        console.log('In Reject')
        myData = Common.getForm(myForm)
        ApplicationAPI.rejectLeaveApplication(myData).then(resp => {
            console.log(resp)
        })
    })
})


// Fill in the update form hidden value with database ID
function showUpdate(id) {
    document.updateLeaveAppForm.ID.value = id
    return true
}