const ReportHelpers = new LeaveReport_Helpers(),
      ReportAPI = new LeaveReport_API(),
      dt = 'leaveReport',
      dtBody = 'leaveReportBody'

window.addEventListener('DOMContentLoaded', () => {
    // fetch all leave types & update DOM
    ReportAPI.getAllLeaveReport().then(resp => {
        console.log(resp)
        ReportHelpers.insertRows(dtBody, resp.data)
        ReportHelpers.triggerDT(dt, dtBody)
    })
})