import datetime
import time
from selenium import webdriver
from selenium.webdriver.common.action_chains import ActionChains
from selenium.webdriver.common.by import By
from selenium.webdriver.support.wait import WebDriverWait
from selenium.webdriver.support import expected_conditions as EC

# t = int(input("请输入播放时间(min): "))


option = webdriver.ChromeOptions()
option.add_argument("--mute-audio")
option.add_argument('--user-data-dir=" + f"/home/west/.config/google-chrome/')
option.add_experimental_option("excludeSwitches", ['enable-automation'])
option.add_experimental_option("useAutomationExtension", False)
browser = webdriver.Chrome(options=option)
browser.maximize_window()
wait = WebDriverWait(browser, 60)
browser.get('https://onlineweb.zhihuishu.com/onlinestuh5')
browser.execute_cdp_cmd("Page.addScriptToEvaluateOnNewDocument", {
    "source": """Object.defineProperty(navigator, 'webdriver', {get: () => undefined})""",
})
USERNAME = '17537011561'
PASSWORD = 'Xcq2540.'
time.sleep(2)
browser.find_element(By.XPATH, value='//*[@id="lUsername"]').send_keys(USERNAME)
time.sleep(0.1)
browser.find_element(By.XPATH, value='//*[@id="lPassword"]').send_keys(PASSWORD)
time.sleep(0.1)
browser.find_element(By.XPATH, value='//*[@id="f_sign_up"]/div[1]/span').click()
time.sleep(5)
wait.until(
    EC.element_to_be_clickable((By.XPATH, '//*[@id="sharingClassed"]/div[2]/ul[2]/div/dl/dt/div[1]/div[1]'))).click()
# wait.until(EC.element_to_be_clickable((By.XPATH, '//*[@id="sharingClassed"]/div[2]/ul[1]/div/dl/dt/div[1]/div[1]'))).click()
time.sleep(5)


# wait.until(EC.element_to_be_clickable((By.XPATH, '//*[@id="playTopic-dialog"]/div/div[6]/div/div[3]/span/button/span'))).click()
# wait.until(EC.element_to_be_clickable((By.CSS_SELECTOR, '#playTopic-dialog > div > div:nth-child(6) > div.dialog-read > div.el-dialog__header')))
# browser.find_element(By.CSS_SELECTOR, value="#playTopic-dialog > div > div:nth-child(6) > div.dialog-read > div.el-dialog__header > i").click()
# time.sleep(3)

def move_mouse():
    """鼠标悬停"""
    ActionChains(browser).move_to_element(
        browser.find_element(By.XPATH, value='/html/body/div[1]/div/div[2]/div[1]/div[2]/div/div/div[8]')).perform()
    time.sleep(0.5)


def star_play():
    """开始播放"""
    move_mouse()
    wait.until(EC.element_to_be_clickable((By.XPATH, '//*[@id="playButton"]'))).click()


def next_video():
    """下一集"""
    move_mouse()
    wait.until(EC.element_to_be_clickable((By.XPATH, '//*[@id="nextBtn"]'))).click()


def fast_speed():
    """倍速播放"""
    move_mouse()
    time.sleep(0.5)
    ActionChains(browser).move_to_element(
        browser.find_element(By.XPATH, value='//*[@id="vjs_container"]/div[10]/div[8]')).perform()
    time.sleep(0.5)
    ActionChains(browser).move_to_element(
        browser.find_element(By.XPATH, value='//*[@id="vjs_container"]/div[10]/div[8]/div/div[1]')).perform()
    ActionChains(browser).click(on_element=None).perform()


def change_play_model():
    """切换流畅模式"""
    move_mouse()
    ActionChains(browser).move_to_element(
        browser.find_element(By.XPATH, value='//*[@id="vjs_container"]/div[10]/div[7]')).perform()
    time.sleep(0.5)
    ActionChains(browser).move_to_element(
        browser.find_element(By.XPATH, value='//*[@id="vjs_container"]/div[10]/div[7]/div/b[1]')).perform()
    ActionChains(browser).click(on_element=None).perform()


def find_window():
    """答题检测"""
    try:
        browser.find_element(By.CSS_SELECTOR, value='#playTopic-dialog > div')
        browser.find_element(By.CSS_SELECTOR,
                             value='#playTopic-dialog > div > div.el-dialog__body > div > div.el-scrollbar__wrap > div > div > div.radio > ul > li:nth-child(1)').click()
        browser.execute_script(
            'window.scrollTo(0,document.body.scrollHeight)'
        )
        try:
            browser.find_element(By.XPATH,
                                    value='//*[@id="playTopic-dialog"]/div/div[2]/div/div[1]/div/div/div[2]/div/p')
            print(browser.find_element(By.XPATH,
                                       value='//*[@id="playTopic-dialog"]/div/div[2]/div/div[1]/div/div/div[2]/div/p').text)
            m = browser.find_element(By.XPATH,
                                     value='//*[@id="playTopic-dialog"]/div/div[2]/div/div[1]/div/div/div[2]/div/p')
            if m.text[-1] == 'A':
                browser.find_element(By.XPATH,
                                     value='//*[@id="playTopic-dialog"]/div/div[2]/div/div[1]/div/div/div[2]/ul/li[1]/span').click()
                browser.find_element(By.XPATH, value='//*[@id="playTopic-dialog"]/div/div[3]/span/div').click()
            elif m.text[-1] == 'B':
                browser.find_element(By.XPATH,
                                     value='//*[@id="playTopic-dialog"]/div/div[2]/div/div[1]/div/div/div[2]/ul/li[2]/span').click()
                browser.find_element(By.XPATH, value='//*[@id="playTopic-dialog"]/div/div[3]/span/div').click()
            elif m.text[-1] == 'C':
                browser.find_element(By.XPATH,
                                     value='//*[@id="playTopic-dialog"]/div/div[2]/div/div[1]/div/div/div[2]/ul/li[3]/span').click()
                browser.find_element(By.XPATH, value='//*[@id="playTopic-dialog"]/div/div[3]/span/div').click()
            elif m.text[-1] == 'D':
                browser.find_element(By.XPATH,
                                     value='//*[@id="playTopic-dialog"]/div/div[2]/div/div[1]/div/div/div[2]/ul/li[4]/span').click()
                browser.find_element(By.XPATH, value='//*[@id="playTopic-dialog"]/div/div[3]/span/div').click()
            else:
                browser.find_element(By.XPATH, value='//*[@id="playTopic-dialog"]/div/div[3]/span/div').click()
        except:
            browser.find_element(By.XPATH,
                                 value='//*[@id="playTopic-dialog"]/div/div[7]/div/div[3]/span/div').click()

    except:
        pass


def find_warn():
    try:
        browser.find_element(By.CSS_SELECTOR,
                                value='#playTopic-dialog > div > div.el-dialog__wrplayTopic-dialoger.dialog-look-habit > div > div.el-dialog__footer > span > button > span')
        browser.find_element(By.CSS_SELECTOR,
                             value='#playTopic-dialog > div > div.el-dialog__wrplayTopic-dialoger.dialog-look-habit > div > div.el-dialog__header > button > i').click()
        time.sleep(0.5)
        star_play()
    except:
        pass


def main(t):
    now = datetime.datetime.now()
    delay = datetime.timedelta(minutes=t)
    new_time = now + delay
    t = new_time.strftime('%Y-%m-%d %H:%M:%S.%f')
    print(f"在{t}时停止")
    while True:
        now = datetime.datetime.now()
        now = now.strftime('%Y-%m-%d %H:%M:%S.%f')
        if now < t:
            try:
                # //*[@id="playTopic-dialog"]/div
                # //*[@id="playTopic-dialog"]/div/div[1]/div/h4
                browser.find_element(By.XPATH, value='//*[@id="playTopic-dialog"]')
                print("检测到答题窗口")
                find_window()
                find_warn()
                time.sleep(0.5)
                star_play()
            except:
                print("未检测到答题窗口")
            try:
                first_time = browser.find_element(By.CSS_SELECTOR,
                                                  value='#vjs_container > div.controlsBar > div.nPlayTime > span.currentTime')
                # print(first_time.text)
                second_time = browser.find_element(By.CSS_SELECTOR,
                                                   value='#vjs_container > div.controlsBar > div.nPlayTime > span.duration')
                # print(second_time.text)
                second_time1 = second_time.text
                if int(second_time.text[3]) == 0:
                    if int(first_time.text[4]) == int(second_time1[4]):
                        print("下一课")
                        next_video()
                        time.sleep(3)
                        move_mouse()
                        star_play()
                        time.sleep(0.5)
                        fast_speed()
                        time.sleep(0.5)
                        change_play_model()
                        time.sleep(0.5)
                elif int(second_time.text[3]) != 0:
                    if int(first_time.text[3:5]) >= int(int(second_time1[3:5]) - 1):
                        print("下一课")
                        next_video()
                        time.sleep(3)
                        move_mouse()
                        star_play()
                        time.sleep(0.5)
                        fast_speed()
                        time.sleep(0.5)
                        change_play_model()
                        time.sleep(0.5)

            except:
                pass
            time.sleep(1)
            move_mouse()
        else:
            browser.close()
            break


# def plan(t):
#     for i in tqdm(range(t*60)):
#         time.sleep(1)
#         yield i


if __name__ == '__main__':
    star_play()
    time.sleep(0.5)
    fast_speed()
    time.sleep(0.5)
    change_play_model()
    time.sleep(0.5)
    count = 0
    time.sleep(5)
    main(30)

