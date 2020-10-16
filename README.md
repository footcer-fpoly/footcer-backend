# footcer-backend
footcer-backend Go - PostgreSQL 
HTTP Status Code : 
- Success :
  + HTTPCode : 200 ( Dữ liệu được gửi lên server đã nhận được)
  + Status Code : 200 -> Server đã xử lí thành công và trả về response.
- Fail:
    + HTTPCode : 200 ( Dữ liệu được gửi lên server đã nhận được)
    + Status Code : 
       - StatusOK (200)                  -> Server đã xử lí thành công và trả về response
       - StatusBadRequest (400)          -> Dữ liệu từ client gửi lên có vấn đề (Thiếu params truyền vào....)
       - StatusConflict (409)            -> Server đã nhận được request nhưng lỗi dữ liệu không tồn tại hoặc câu Query trên server có vấn đề....
       - StatusInternalServerError (500) -> Server gen token thất bại hoặc server CHẾT. =)))
    
