import requests

r = requests.get("http://stronikum.ru/4409_Vtoraya_mladshaya_gruppa_3_4")
res = r.content.decode("utf-8")

print(str(res))

# for test stronikum
# result html part*:

        #  <tr class="product-row">
        #     <td><a href="4409_Vtoraya_mladshaya_gruppa_3_4/3769_Buben">Бубен</a></td>
        #     <td width="10%">
        #       <input class="price-count" type="number" min="1" value="1" disabled>
        #       <input type="checkbox">
        #     </td>
        #     <td width="10%" align="right">1950 р.</td>
        #   </tr>
        #           <tr class="product-row">
        #     <td><a href="4409_Vtoraya_mladshaya_gruppa_3_4/11560_Komplekt_shumovih_muzikalnih_instrumentov">Комплект шумовых музыкальных инструментов</a></td>
        #     <td width="10%">
        #       <input class="price-count" type="number" min="1" value="1" disabled>
        #       <input type="checkbox">
        #     </td>
        #     <td width="10%" align="right">5880 р.</td>
        #   </tr>
        #           <tr class="product-row">
        #     <td><a href="4409_Vtoraya_mladshaya_gruppa_3_4/4098_Logki_para">Ложки (пара)</a></td>
        #     <td width="10%">
        #       <input class="price-count" type="number" min="1" value="1" disabled>
        #       <input type="checkbox">
        #     </td>
        #     <td width="10%" align="right">650 р.</td>
        #   </tr>
        #           <tr class="product-row">
        #     <td><a href="4409_Vtoraya_mladshaya_gruppa_3_4/4100_Marakas">Маракас</a></td>
        #     <td width="10%">
        #       <input class="price-count" type="number" min="1" value="1" disabled>
        #       <input type="checkbox">
        #     </td>
        #     <td width="10%" align="right">1700 р.</td>
        #   </tr>