import React from "react";
import Link from "next/link";
import { Button } from "@/components/ui/button";
import { routesAuth } from "@/configs/routes";
import styles from "./styles.module.scss";

const StartPage: React.FC = () => {
    return <main className={ styles.main }>
        <div className={ styles.content }>
            <h3 className={ styles.title }>Faculty allocation helper</h3>
            <div className={ styles.description }>To get access to the service, please, login through your Innopolis University e-mail. </div>
            <div className={ styles.buttons }>
                {
                    routesAuth.map(({ routeName, routePath }) =>
                      <Button key={ routePath } variant={ "strictWhite" } className={ styles.button }>
                          <Link href={ routePath }>{ routeName }</Link>
                      </Button>)
                }
            </div>
        </div>
    </main>;
};

export default StartPage;