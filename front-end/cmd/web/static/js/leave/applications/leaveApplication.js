const  Helpers = new LeaveApplication_Helpers(),
       API = new LeaveApplication_API()

// Set datatable parameters (JQuery Datatable)
const dt = 'leaveApplication',
    dtBody = 'leaveApplicationBody'


window.addEventListener('DOMContentLoaded', () => {
       // fetch all employee summary & update DOM
       API.getAllLeaveApplication().then(resp => { 
        console.log(resp)
        Helpers.insertRows(dtBody, resp.data)
        Helpers.triggerDT(dt, dtBody)
    })
})