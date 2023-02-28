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
})